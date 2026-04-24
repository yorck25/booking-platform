import { useEffect, useState } from "preact/hooks";
import { Select } from "../../components/select";
import Flatpickr from "react-flatpickr";
import "flatpickr/dist/themes/light.css";
import styles from "./style.module.scss";
import { Divider } from "../../components/divider";
import { BookingService } from "../../services/booking-service";

export interface IReservationSlot {
    startTime: string;
    endTime: string;
    isBooked: boolean;
}

export const ReservationPage = () => {
    const [date, setDate] = useState<Date | null>(new Date());
    const [slots, setSlots] = useState<IReservationSlot[]>([]);
    const [time, setTime] = useState("");

    const loadBookingSlots = async (date: Date) => {
        try {
            const reservationSlots = await BookingService.loadBarberAvailability(date);

            setSlots(reservationSlots);

            const firstAvailableSlot = reservationSlots.find(slot => !slot.isBooked);

            if (firstAvailableSlot) {
                setTime(firstAvailableSlot.startTime);
            } else {
                setTime("");
            }
        } catch (error) {
            console.error("Failed to load booking slots", error);
            setSlots([]);
            setTime("");
        }
    };

    useEffect(() => {
        loadBookingSlots(new Date());
    }, []);

    const handleDateChange = (selectedDates: Date[]) => {
        const selectedDate = selectedDates.length > 0 ? selectedDates[0] : null;

        setDate(selectedDate);

        if (selectedDate) {
            loadBookingSlots(selectedDate);
        }
    };

    const handleTimeInput = (e: Event) => {
        const target = e.target as HTMLSelectElement;
        setTime(target.value);
    };

    const getTimeOptions = () => {
        return slots
            .filter(slot => !slot.isBooked)
            .map(slot => ({
                value: slot.startTime,
                label: slot.startTime,
            }));
    };

    const handleSubmit = (e: Event) => {
        e.preventDefault();

        console.log({
            date,
            time,
        });
    };

    return (
        <div className={styles.time_booking_view}>
            <div className={styles.header}>
                <h1>Termin Buchen</h1>
                <Divider />
            </div>

            <form className={styles.reservation_form} onSubmit={handleSubmit}>
                <div className={styles.input_group}>
                    <div className={styles.date_picker_group}>
                        <label htmlFor="reservation-date" className={styles.label}>
                            Datum
                        </label>

                        <Flatpickr
                            value={date || undefined}
                            onChange={handleDateChange}
                            options={{
                                dateFormat: "d-m-Y",
                                allowInput: false,
                            }}
                            className={styles.date_input}
                        />
                    </div>

                    <Select
                        id="time"
                        label="Uhrzeit"
                        value={time}
                        placeholder="Uhrzeit wählen"
                        handleChange={handleTimeInput}
                        options={getTimeOptions()}
                    />
                </div>

                <Divider />

                <div className={styles.suggestion_group}>
                    {slots
                        .filter(slot => !slot.isBooked)
                        .map(slot => (
                            <button
                                key={slot.startTime}
                                type="button"
                                className={styles.suggestion_cell}
                                onClick={() => setTime(slot.startTime)}
                            >
                                <p>{slot.startTime}</p>
                            </button>
                        ))}
                </div>

                <div className={styles.footer_button_wrapper}>
                    <button className={styles.button} type="submit" disabled={!time}>
                        Weiter
                    </button>
                </div>
            </form>
        </div>
    );
};