-- =========================
-- BARBER
-- =========================
INSERT INTO barbers (
    id, internal_name, display_name, email, phone, mobile_phone,
    postal_code, street, city
) VALUES (
             '11111111-1111-1111-1111-111111111111',
             'main_barber_shop',
             'Yorck Barber Shop',
             'info@yorck-barber.de',
             '069123456',
             '+491701234567',
             '60311',
             'Zeil 1',
             'Frankfurt'
         );

-- =========================
-- EMPLOYEE (OWNER)
-- =========================
INSERT INTO barber_employees (
    id, barber_id,
    username, first_name, last_name, display_name, internal_name,
    email, phone,
    password_hash, role
) VALUES (
             '22222222-2222-2222-2222-222222222222',
             '11111111-1111-1111-1111-111111111111',
             'yorck',
             'Yorck',
             'Dombrowsky',
             'Yorck',
             'yorck_owner',
             'yorck@barber.de',
             '+491701234567',
             '$2a$10$dummyhash', -- replace with real bcrypt
             'owner'
         );

-- =========================
-- WORKING HOURS (Mon-Fri)
-- =========================
INSERT INTO employee_working_hours (
    id, employee_id, weekday, start_time, end_time
) VALUES
      ('30000000-0000-0000-0000-000000000001', '22222222-2222-2222-2222-222222222222', 1, '09:00', '18:00'),
      ('30000000-0000-0000-0000-000000000002', '22222222-2222-2222-2222-222222222222', 2, '09:00', '18:00'),
      ('30000000-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', 3, '09:00', '18:00'),
      ('30000000-0000-0000-0000-000000000004', '22222222-2222-2222-2222-222222222222', 4, '09:00', '18:00'),
      ('30000000-0000-0000-0000-000000000005', '22222222-2222-2222-2222-222222222222', 5, '09:00', '18:00');

-- =========================
-- BREAKS (Lunch)
-- =========================
INSERT INTO employee_breaks (
    id, employee_id, weekday, start_time, end_time, description
) VALUES
      ('40000000-0000-0000-0000-000000000001', '22222222-2222-2222-2222-222222222222', 1, '13:00', '13:30', 'Lunch'),
      ('40000000-0000-0000-0000-000000000002', '22222222-2222-2222-2222-222222222222', 2, '13:00', '13:30', 'Lunch'),
      ('40000000-0000-0000-0000-000000000003', '22222222-2222-2222-2222-222222222222', 3, '13:00', '13:30', 'Lunch'),
      ('40000000-0000-0000-0000-000000000004', '22222222-2222-2222-2222-222222222222', 4, '13:00', '13:30', 'Lunch'),
      ('40000000-0000-0000-0000-000000000005', '22222222-2222-2222-2222-222222222222', 5, '13:00', '13:30', 'Lunch');

-- =========================
-- CLOSED DAY
-- =========================
INSERT INTO barber_closed_days (
    id, barber_id, closed_date, reason
) VALUES (
             '50000000-0000-0000-0000-000000000001',
             '11111111-1111-1111-1111-111111111111',
             '2026-12-25',
             'Christmas'
         );

-- =========================
-- SERVICES
-- =========================
INSERT INTO services (
    id, barber_id, internal_name, display_name,
    description, duration_minutes, price_cents, sort_order
) VALUES
      (
          '60000000-0000-0000-0000-000000000001',
          '11111111-1111-1111-1111-111111111111',
          'mens_haircut',
          'Herren Haarschnitt',
          'Classic men haircut',
          30,
          2500,
          1
      ),
      (
          '60000000-0000-0000-0000-000000000002',
          '11111111-1111-1111-1111-111111111111',
          'beard_trim',
          'Bart trimmen',
          'Beard styling and trim',
          20,
          1500,
          2
      );

-- =========================
-- EMPLOYEE SERVICES
-- =========================
INSERT INTO employee_services (
    id, employee_id, service_id
) VALUES
      ('70000000-0000-0000-0000-000000000001', '22222222-2222-2222-2222-222222222222', '60000000-0000-0000-0000-000000000001'),
      ('70000000-0000-0000-0000-000000000002', '22222222-2222-2222-2222-222222222222', '60000000-0000-0000-0000-000000000002');

-- =========================
-- BOOKING
-- =========================
INSERT INTO bookings (
    id, barber_id, employee_id, service_id,
    customer_first_name, customer_last_name, customer_phone_number, customer_email,
    booking_date, start_time, end_time,
    status,
    service_name, service_duration_min, service_price_cents,
    terms_accepted, terms_accepted_at
) VALUES (
             '80000000-0000-0000-0000-000000000001',
             '11111111-1111-1111-1111-111111111111',
             '22222222-2222-2222-2222-222222222222',
             '60000000-0000-0000-0000-000000000001',
             'Max',
             'Mustermann',
             '+491701112233',
             'max@example.com',
             CURRENT_DATE,
             '10:00',
             '10:30',
             'confirmed',
             'Herren Haarschnitt',
             30,
             2500,
             true,
             now()
         );

-- =========================
-- BOOKING EVENT
-- =========================
INSERT INTO booking_events (
    id, booking_id, event_type, message
) VALUES (
             '90000000-0000-0000-0000-000000000001',
             '80000000-0000-0000-0000-000000000001',
             'created',
             'Booking created successfully'
         );