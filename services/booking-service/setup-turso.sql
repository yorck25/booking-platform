PRAGMA foreign_keys = ON;

CREATE TABLE barbers (
  id TEXT PRIMARY KEY,
  internal_name TEXT NOT NULL,
  display_name TEXT NOT NULL,
  email TEXT,
  phone TEXT,
  mobile_phone TEXT,
  postal_code TEXT,
  street TEXT,
  city TEXT,
  active INTEGER NOT NULL DEFAULT 1,
  deleted INTEGER NOT NULL DEFAULT 0,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE barber_employees (
  id TEXT PRIMARY KEY,
  barber_id TEXT NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
  username TEXT NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  display_name TEXT NOT NULL,
  internal_name TEXT NOT NULL,
  email TEXT,
  phone TEXT,
  password_hash TEXT NOT NULL,
  role TEXT NOT NULL DEFAULT 'user',
  last_login TEXT,
  failed_attempts INTEGER NOT NULL DEFAULT 0,
  active INTEGER NOT NULL DEFAULT 1,
  deleted INTEGER NOT NULL DEFAULT 0,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT chk_barber_employees_role CHECK (role IN ('owner', 'admin', 'user')),
  CONSTRAINT chk_barber_employees_failed_attempts CHECK (failed_attempts >= 0),
  CONSTRAINT uq_barber_employees_username UNIQUE (username)
);

CREATE TABLE employee_working_hours (
  id TEXT PRIMARY KEY,
  employee_id TEXT NOT NULL REFERENCES barber_employees (id) ON DELETE CASCADE,
  weekday INTEGER NOT NULL CHECK (weekday BETWEEN 0 AND 6),
  start_time TEXT NOT NULL,
  end_time TEXT NOT NULL,
  is_closed INTEGER NOT NULL DEFAULT 0,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT chk_employee_working_hours_time CHECK (
    is_closed = 1
    OR (
      is_closed = 0
      AND end_time > start_time
    )
  )
);

CREATE TABLE employee_breaks (
  id TEXT PRIMARY KEY,
  employee_id TEXT NOT NULL REFERENCES barber_employees (id) ON DELETE CASCADE,
  weekday INTEGER NOT NULL CHECK (weekday BETWEEN 0 AND 6),
  start_time TEXT NOT NULL,
  end_time TEXT NOT NULL,
  description TEXT,
  active INTEGER NOT NULL DEFAULT 1,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT chk_employee_breaks_time CHECK (end_time > start_time)
);

CREATE TABLE barber_closed_days (
  id TEXT PRIMARY KEY,
  barber_id TEXT NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
  closed_date TEXT NOT NULL,
  reason TEXT,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT uq_barber_closed_days UNIQUE (barber_id, closed_date)
);

CREATE TABLE services (
  id TEXT PRIMARY KEY,
  barber_id TEXT NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
  internal_name TEXT NOT NULL,
  display_name TEXT NOT NULL,
  description TEXT,
  category TEXT NOT NULL DEFAULT 'other',
  duration_minutes INTEGER NOT NULL,
  price_cents INTEGER NOT NULL,
  active INTEGER NOT NULL DEFAULT 1,
  deleted INTEGER NOT NULL DEFAULT 0,
  sort_order INTEGER NOT NULL DEFAULT 0,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT chk_services_duration CHECK (duration_minutes > 0),
  CONSTRAINT chk_services_price CHECK (price_cents >= 0)
);

CREATE TABLE employee_services (
  id TEXT PRIMARY KEY,
  employee_id TEXT NOT NULL REFERENCES barber_employees (id) ON DELETE CASCADE,
  service_id TEXT NOT NULL REFERENCES services (id) ON DELETE CASCADE,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT uq_employee_services UNIQUE (employee_id, service_id)
);

CREATE TABLE bookings (
  id TEXT PRIMARY KEY,
  barber_id TEXT NOT NULL REFERENCES barbers (id) ON DELETE CASCADE,
  employee_id TEXT NOT NULL REFERENCES barber_employees (id) ON DELETE RESTRICT,
  service_id TEXT NOT NULL REFERENCES services (id) ON DELETE RESTRICT,
  customer_first_name TEXT NOT NULL,
  customer_last_name TEXT NOT NULL,
  customer_phone_number TEXT NOT NULL,
  customer_email TEXT,
  booking_date TEXT NOT NULL,
  start_time TEXT NOT NULL,
  end_time TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'pending',
  cancel_reason TEXT,
  notes TEXT,
  service_name TEXT NOT NULL,
  service_duration_min INTEGER NOT NULL,
  service_price_cents INTEGER NOT NULL,
  terms_accepted INTEGER NOT NULL DEFAULT 0,
  terms_accepted_at TEXT,
  confirmed_at TEXT,
  rejected_at TEXT,
  canceled_at TEXT,
  sms_sent INTEGER NOT NULL DEFAULT 0,
  sms_sent_at TEXT,
  sms_error TEXT,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT chk_bookings_time CHECK (end_time > start_time),
  CONSTRAINT chk_bookings_status CHECK (
    status IN ('pending', 'confirmed', 'rejected', 'cancelled')
  ),
  CONSTRAINT chk_bookings_service_duration CHECK (service_duration_min > 0),
  CONSTRAINT chk_bookings_service_price CHECK (service_price_cents >= 0)
);

CREATE TABLE booking_events (
  id TEXT PRIMARY KEY,
  booking_id TEXT NOT NULL REFERENCES bookings (id) ON DELETE CASCADE,
  event_type TEXT NOT NULL,
  message TEXT,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
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