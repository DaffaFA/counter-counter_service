package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/DaffaFA/counter-counter_service/utils"
	"github.com/bytedance/sonic"
	"github.com/exaring/otelpgx"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ctx := context.Background()

	db := getDB(ctx)
	defer db.Close()

	rdb := getRDB()
	defer rdb.Close()

	otelConn, err := initConn()
	if err != nil {
		log.Panic().AnErr("error", err).Msg("failed to create gRPC connection to collector")
	}

	_, err = utils.InitTracerProvider(ctx, resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(viper.GetString("SERVICE_NAME")),
	), otelConn)
	if err != nil {
		log.Panic().AnErr("error", err).Msg("failed to create trace exporter")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		Prefork:     false})

	if viper.GetBool("DEBUG") {
		app.Use(logger.New())
	}

	app.Use(otelfiber.Middleware())

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/__health",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			pctx, pctxCancel := context.WithDeadline(c.Context(), time.Now().Add(5*time.Second))
			defer pctxCancel()

			err := db.Ping(pctx)

			return err == nil
		},
		ReadinessEndpoint: "/__ready",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	app.Get("/__monitor", monitor.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// urlPrefix := app.Group("/api/v1")

	app.Listen(fmt.Sprintf(":%s", viper.GetString("PORT")))
}

func getDB(ctx context.Context) *pgxpool.Pool {
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASS")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	log.Info().Str("dbUser", dbUser).Str("dbHost", dbHost).Str("dbName", dbName).Msg("connecting to database")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	cfg, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Panic().AnErr("error", err).Msg("failed to connect to database")
	}

	cfg.ConnConfig.Tracer = otelpgx.NewTracer()

	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Panic().AnErr("error", err).Msg("failed to connect to database")
	}
	return db
}

func getRDB() *redis.Client {
	rdbPort := viper.GetString("REDIS_PORT")
	rdbHost := viper.GetString("REDIS_HOST")
	rdbPassword := viper.GetString("REDIS_PASS")
	rdbName := viper.GetInt("REDIS_DB")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdbHost, rdbPort),
		Password: rdbPassword,
		DB:       rdbName,
	})

	status := rdb.Ping(ctx)

	if status.Err() != nil {
		log.Panic().AnErr("error", status.Err()).Msg("failed to connect to redis")
	}

	log.Info().Msg("connected to redis")

	return rdb

}

func initConn() (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}
