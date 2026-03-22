import { useState } from "preact/hooks";
import { Select } from "../../components/select";
import Flatpickr from "react-flatpickr";
import "flatpickr/dist/themes/light.css";
import styles from "./style.module.scss";
import { Divider } from "../../components/divider";

export const ReservationPage = () => {
    const times = [
        "09:00",
        "09:30",
        "10:00",
        "10:30",
    ];

    const [date, setDate] = useState<Date | null>(null);
    const [time, setTime] = useState(times[0]);

    const handleTimeInput = (e: Event) => {
        const target = e.target as HTMLSelectElement;
        setTime(target.value);
    };

    const getTimeOptions = () => {
        return times.map((t) => ({
            value: t,
            label: t,
        }));
    };

    const handleSubmit = (e: Event) => {
        e.preventDefault();
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
                            onChange={(selectedDates) => {
                                setDate(selectedDates.length > 0 ? selectedDates[0] : null);
                            }}
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
                        placeholder=""
                        handleChange={handleTimeInput}
                        options={getTimeOptions()}
                    />
                </div>

                <Divider />

                <div className={styles.suggestion_group}>
                    <div className={styles.suggestion_cell}><p>08:30</p></div>
                    <div className={styles.suggestion_cell}><p>09:00</p></div>
                    <div className={styles.suggestion_cell}><p>09:30</p></div>
                    <div className={styles.suggestion_cell}><p>10:00</p></div>
                    <div className={styles.suggestion_cell}><p>10:30</p></div>
                    <div className={styles.suggestion_cell}><p>11:00</p></div>
                    <div className={styles.suggestion_cell}><p>11:30</p></div>
                    <div className={styles.suggestion_cell}><p>12:00</p></div>
                </div>

                <div className={styles.footer_button_wrapper}>
                    <button className={styles.button} type="submit">
                        Weiter
                    </button>
                </div>
            </form>
        </div>
    );
};