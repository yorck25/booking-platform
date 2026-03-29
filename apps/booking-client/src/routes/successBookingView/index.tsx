import styles from "./style.module.scss";
import {CheckIcon, CircleInfo} from "../../components/icons";
import {ReservationDetails} from "../../components/reservationDetails";

export const SuccessBookingView = () => {

    const reservationDetailsProps = {
        activeConfig: {
            showService: true,
            showDate: true,
            showTime: true,
            showName: true,
            showPhoneNumer: true,
        }
    }

    return (
        <div className={styles.success_booking_view}>
            <div className={styles.check_icon_container_wrapper}>
                <div className={styles.check_icon_container}>
                    {CheckIcon()}
                </div>
            </div>

            <div className={styles.description}>
                <h1>Deine Buchung wurde verschickt</h1>

                <div className={styles.description_text}>
                    <p className={styles.row}>Dein Friseur muss den Termin noch besätigen</p>
                    <p className={styles.row}>Sobald der Termin bestätigt wurde, erhälst du eine SMS an die angegebene
                        Telefonnummer</p>
                </div>
            </div>

            <div className={styles.details}>
                <ReservationDetails activeConfig={reservationDetailsProps.activeConfig}/>
            </div>

            <div className={styles.hint}>
                <div className={styles.leading_icon}>
                    {CircleInfo()}
                </div>

                <div className={styles.hint_description}>
                    <p className={styles.hint_headline}>Hinweis</p>
                    <p className={styles.hint_text}>Die SMS wird and die angegebene Telefonnummer geschickt.</p>
                </div>
            </div>

            <div className={styles.footer_button_wrapper}>
                <button className={styles.button} type="submit">
                    Fertig
                </button>
            </div>
        </div>
    )
}