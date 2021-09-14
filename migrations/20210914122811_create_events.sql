-- +goose Up
-- +goose StatementBegin
CREATE TABLE events
(
  ip            IPv4,
  server_time   DateTime,
  client_time   DateTime,
  device_id     UUID,
  device_os     String,
  session       String,
  sequence      Int8,
  event         String,
  param_int     Int8,
  param_str     String
) ENGINE = MergeTree()
ORDER BY (server_time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
