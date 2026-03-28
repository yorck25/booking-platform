import styles from "./style.module.scss";
import {Divider} from "../../components/divider";
import {useState} from "preact/hooks";
import {Input, InputType} from "../../components/input";
import {ReservationDetails} from "../../components/reservationDetails";

export const PersonalDataPage = () => {
    // const date = "12.01.2026";
    // const time = "09:30";

    const [firstName, setFristName] = useState<string>("");
    const [lastName, setLastName] = useState<string>("");
    const [phoneNumer, setPhoneNumer] = useState<string>("");


    const handleFirstNameInput = (e: Event) => {
        const target = e.target as HTMLSelectElement;
        setFristName(target.value);
    };

    const handleLastNameInput = (e: Event) => {
        const target = e.target as HTMLSelectElement;
        setLastName(target.value);
    };

    const handlePhoneInput = (e: Event) => {
        const target = e.target as HTMLSelectElement;
        setPhoneNumer(target.value);
    };


    const handleSubmit = (e: Event) => {
        e.preventDefault();
    };

    const reservationDetailsProps = {
        activeConfig: {
            showService: true,
            showDate: true,
            showTime: true,
            showName: false,
            showPhoneNumer: false,
        }
    }

    return (
        <div className={styles.time_booking_view}>
            <div className={styles.header}>
                <div className={styles.header_row}>
                    <h1>Persönliche Daten</h1>
                </div>

                <Divider/>
            </div>

            <div className={styles.details_render_container}>
                <div className={styles.details}>
                    <ReservationDetails activeConfig={reservationDetailsProps.activeConfig}/>
                </div>

                <Divider/>
            </div>

            <form className={styles.personal_data_form} onSubmit={handleSubmit}>
                <div className={styles.input_group}>
                    <div className={styles.form_row}>
                        <Input
                            label={"Vorname"}
                            inputType={InputType.TEXT}
                            id="firstName"
                            placeholder="Max"
                            value={firstName}
                            handleInput={handleFirstNameInput}
                        />

                        <Input
                            label={"Nachname"}
                            inputType={InputType.TEXT}
                            id="lastName"
                            placeholder="Musterman"
                            value={lastName}
                            handleInput={handleLastNameInput}
                        />
                    </div>

                    <Input
                        label={"Telefonnummer"}
                        inputType={InputType.TEXT}
                        id="phoneNumer"
                        placeholder="+49 151 12345678"
                        value={phoneNumer}
                        handleInput={handlePhoneInput}
                    />
                </div>


                <div className={styles.footer_button_wrapper}>
                    <button className={styles.button} type="submit">
                        Buchen
                    </button>
                </div>
            </form>
        </div>
    );
};