CREATE TABLE barbers
(
    id            uuid PRIMARY KEY,
    internal_name varchar(255) NOT NULL,
    display_name  varchar(255) NOT NULL,
    email         varchar(255),
    phone         varchar(50),
    mobile_phone  varchar(50),
    postal_code   varchar(20),
    street        varchar(128),
    city          varchar(128),
    active        boolean      NOT NULL DEFAULT true,
    deleted       boolean      NOT NULL DEFAULT false,
    created_at    timestamp    NOT NULL DEFAULT now(),
    updated_at    timestamp    NOT NULL DEFAULT now()
);

CREATE TABLE barber_employees
(
    id              uuid PRIMARY KEY,
    barber_id       uuid         NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,

    username        varchar(128) NOT NULL,
    first_name      varchar(128) NOT NULL,
    last_name       varchar(128) NOT NULL,
    display_name    varchar(255) NOT NULL,
    internal_name   varchar(255) NOT NULL,

    email           varchar(255),
    phone           varchar(50),

    password_hash   varchar(255) NOT NULL,
    role            varchar(32)  NOT NULL DEFAULT 'user',

    last_login      timestamp,
    failed_attempts smallint     NOT NULL DEFAULT 0,

    active          boolean      NOT NULL DEFAULT true,
    deleted         boolean      NOT NULL DEFAULT false,

    created_at      timestamp    NOT NULL DEFAULT now(),
    updated_at      timestamp    NOT NULL DEFAULT now(),

    CONSTRAINT chk_barber_employees_role
        CHECK (role IN ('owner', 'admin', 'user')),

    CONSTRAINT chk_barber_employees_failed_attempts
        CHECK (failed_attempts >= 0),

    CONSTRAINT uq_barber_employees_username UNIQUE (username)
);

CREATE TABLE employee_working_hours
(
    id          uuid PRIMARY KEY,
    employee_id uuid      NOT NULL REFERENCES barber_employees (id) ON DELETE CASCADE,
    weekday     smallint  NOT NULL CHECK (weekday BETWEEN 0 AND 6),
    start_time  time      NOT NULL,
    end_time    time      NOT NULL,
    is_closed   boolean   NOT NULL DEFAULT false,
    created_at  timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp NOT NULL DEFAULT now(),

    CONSTRAINT chk_employee_working_hours_time
        CHECK (
            (is_closed = true)
                OR
            (is_closed = false AND end_time > start_time)
            )
);

CREATE TABLE employee_breaks
(
    id          uuid PRIMARY KEY,
    employee_id uuid      NOT NULL REFERENCES barber_employees (id) ON DELETE CASCADE,
    weekday     smallint  NOT NULL CHECK (weekday BETWEEN 0 AND 6),
    start_time  time      NOT NULL,
    end_time    time      NOT NULL,
    description varchar(255),
    active      boolean   NOT NULL DEFAULT true,
    created_at  timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp NOT NULL DEFAULT now(),

    CONSTRAINT chk_employee_breaks_time
        CHECK (end_time > start_time)
);

CREATE TABLE barber_closed_days
(
    id          uuid PRIMARY KEY,
    barber_id   uuid         NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
    closed_date date         NOT NULL,
    reason      varchar(255),
    created_at  timestamp    NOT NULL DEFAULT now(),

    CONSTRAINT uq_barber_closed_days UNIQUE (barber_id, closed_date)
);

CREATE TABLE services
(
    id               uuid PRIMARY KEY,
    barber_id        uuid         NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
    internal_name    varchar(255) NOT NULL,
    display_name     varchar(255) NOT NULL,
    description      text,
    duration_minutes integer      NOT NULL,
    price_cents      integer      NOT NULL,
    active           boolean      NOT NULL DEFAULT true,
    deleted          boolean      NOT NULL DEFAULT false,
    sort_order       integer      NOT NULL DEFAULT 0,
    created_at       timestamp    NOT NULL DEFAULT now(),
    updated_at       timestamp    NOT NULL DEFAULT now(),

    CONSTRAINT chk_services_duration
        CHECK (duration_minutes > 0),

    CONSTRAINT chk_services_price
        CHECK (price_cents >= 0)
);

CREATE TABLE employee_services
(
    id          uuid PRIMARY KEY,
    employee_id uuid      NOT NULL REFERENCES barber_employees (id) ON DELETE CASCADE,
    service_id  uuid      NOT NULL REFERENCES services (id) ON DELETE CASCADE,
    created_at  timestamp NOT NULL DEFAULT now(),

    CONSTRAINT uq_employee_services UNIQUE (employee_id, service_id)
);

CREATE TABLE bookings
(
    id                    uuid PRIMARY KEY,
    barber_id             uuid         NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
    employee_id           uuid         NOT NULL REFERENCES barber_employees (id) ON DELETE RESTRICT,
    service_id            uuid         NOT NULL REFERENCES services (id) ON DELETE RESTRICT,

    customer_first_name   varchar(128) NOT NULL,
    customer_last_name    varchar(255) NOT NULL,
    customer_phone_number varchar(50)  NOT NULL,
    customer_email        varchar(255),

    booking_date          date         NOT NULL,
    start_time            time         NOT NULL,
    end_time              time         NOT NULL,

    status                varchar(32)  NOT NULL DEFAULT 'pending',
    cancel_reason         varchar(255),
    notes                 text,

    service_name          varchar(255) NOT NULL,
    service_duration_min  integer      NOT NULL,
    service_price_cents   integer      NOT NULL,

    terms_accepted        boolean      NOT NULL DEFAULT false,
    terms_accepted_at     timestamp,

    confirmed_at          timestamp,
    rejected_at           timestamp,
    canceled_at           timestamp,

    sms_sent              boolean      NOT NULL DEFAULT false,
    sms_sent_at           timestamp,
    sms_error             text,

    created_at            timestamp    NOT NULL DEFAULT now(),
    updated_at            timestamp    NOT NULL DEFAULT now(),

    CONSTRAINT chk_bookings_time
        CHECK (end_time > start_time),

    CONSTRAINT chk_bookings_status
        CHECK (status IN ('pending', 'confirmed', 'rejected', 'cancelled')),

    CONSTRAINT chk_bookings_service_duration
        CHECK (service_duration_min > 0),

    CONSTRAINT chk_bookings_service_price
        CHECK (service_price_cents >= 0)
);

CREATE TABLE booking_events
(
    id          uuid PRIMARY KEY,
    booking_id  uuid         NOT NULL REFERENCES bookings (id) ON DELETE CASCADE,
    event_type  varchar(50)  NOT NULL,
    message     text,
    created_at  timestamp    NOT NULL DEFAULT now()
);

CREATE INDEX idx_barber_employees_barber_id
    ON barber_employees (barber_id);

CREATE INDEX idx_barber_employees_barber_active
    ON barber_employees (barber_id, active, deleted);

CREATE INDEX idx_barber_employees_email
    ON barber_employees (email);

CREATE INDEX idx_employee_working_hours_employee_weekday
    ON employee_working_hours (employee_id, weekday);

CREATE INDEX idx_employee_breaks_employee_weekday
    ON employee_breaks (employee_id, weekday);

CREATE INDEX idx_barber_closed_days_barber_date
    ON barber_closed_days (barber_id, closed_date);

CREATE INDEX idx_services_barber_id
    ON services (barber_id);

CREATE INDEX idx_services_barber_active
    ON services (barber_id, active, deleted);

CREATE INDEX idx_employee_services_employee_id
    ON employee_services (employee_id);

CREATE INDEX idx_employee_services_service_id
    ON employee_services (service_id);

CREATE INDEX idx_bookings_barber_date
    ON bookings (barber_id, booking_date);

CREATE INDEX idx_bookings_employee_date
    ON bookings (employee_id, booking_date);

CREATE INDEX idx_bookings_employee_date_status
    ON bookings (employee_id, booking_date, status);

CREATE INDEX idx_bookings_status
    ON bookings (status);

CREATE INDEX idx_booking_events_booking_id
    ON booking_events (booking_id);
