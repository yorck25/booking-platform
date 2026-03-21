import {useState} from "preact/hooks";
import {Select} from "../../components/select";
import Flatpickr from "react-flatpickr";

import("flatpickr/dist/themes/light.css");
import styles from "./style.module.scss";
import {Divider} from "../../components/divider";

export const ReservationPage = () => {
    const times = [
        "09:00",
        "09:30",
        "10:00",
        "10:30",
    ];

    const [date, setDate] = useState("");
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

    return (
        <div className={styles.time_booking_view}>
            <div className={styles.header}>
                <h1>Termin Buchen</h1>

                <Divider />
            </div>

            <form className={styles.reservation_form}>
                <div className={styles.input_group}>
                    <div className={styles.date_picker_group}>
                        <label htmlFor="reservation-date" className={styles.label}>
                            Datum
                        </label>

                        <Flatpickr
                            value={date}
                            onChange={(_, dateStr) => setDate(dateStr)}
                            options={{
                                dateFormat: "d-m-y",
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
                    {/*Todo: Replace with render function*/}
                    <div className={styles.suggestion_cell}>
                        <p>08:30</p>
                    </div>

                    <div className={styles.suggestion_cell}>
                        <p>09:00</p>
                    </div>

                    <div className={styles.suggestion_cell}>
                        <p>09:30</p>
                    </div>

                    <div className={styles.suggestion_cell}>
                        <p>10:00</p>
                    </div>

                    <div className={styles.suggestion_cell}>
                        <p>10:30</p>
                    </div>

                    <div className={styles.suggestion_cell}>
                        <p>11:00</p>
                    </div>


                    <div className={styles.suggestion_cell}>
                        <p>11:30</p>
                    </div>

                    <div className={styles.suggestion_cell}>
                        <p>12:00</p>
                    </div>
                </div>

                <div className={styles.footer_button_wrapper}>
                    <button className={styles.button} type={"submit"}>Weiter</button>
                </div>
            </form>
        </div>
    );
};
