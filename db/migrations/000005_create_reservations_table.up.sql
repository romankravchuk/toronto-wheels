CREATE TABLE
    IF NOT EXISTS reservations (
        id uuid NOT NULL,
        customer_id uuid NOT NULL,
        vehicle_id uuid NOT NULL,
        start_date timestamptz NOT NULL,
        end_date timestamptz NOT NULL,
        total_cost money NOT NULL,
        pickup_location_id uuid NOT NULL,
        dropoff_location_id uuid NOT NULL,
        status integer NOT NULL,
        CONSTRAINT pk_reservations PRIMARY KEY (id),
        CONSTRAINT unq_reservations_customer_id UNIQUE (customer_id),
        CONSTRAINT unq_reservations_vehicle_id UNIQUE (vehicle_id),
        CONSTRAINT unq_reservations_pickup_location_id UNIQUE (pickup_location_id),
        CONSTRAINT unq_reservations_dropoff_location_id UNIQUE (dropoff_location_id)
    );

ALTER TABLE vehicles
ADD
    CONSTRAINT fk_vehicles_reservations FOREIGN KEY (id) REFERENCES reservations(vehicle_id);