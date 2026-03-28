import styles from "./style.module.scss";
import {CalendarIcon, ClockIcon, PhoneIcon, ScissorsIcon, UserIcon} from "../icons";

interface IProps {
    activeConfig: {
        showService: boolean,
        showDate: boolean,
        showTime: boolean,
        showName: boolean,
        showPhoneNumer: boolean,
    }
}

export const ReservationDetails = (props: IProps) => {
    return (
        <div className={styles.reservation_details_container}>
            <div className={styles.header}>
                <div className={styles.header_icon}>{CalendarIcon()}</div>
                <p className={styles.header_title}>Deine Termin-Details</p>
            </div>

            <ul className={styles.list}>
                {props.activeConfig.showService && (
                    <li className={styles.list_item}>
                        <div className={styles.list_item_left}>
                            <div className={styles.list_item_icon}>{ScissorsIcon()}</div>
                            <p className={styles.list_item_label}>Service</p>
                        </div>

                        <p className={styles.list_item_value}>Frauen Kurzharr Schnitt</p>
                    </li>
                )}

                {props.activeConfig.showDate && (
                    <li className={styles.list_item}>
                        <div className={styles.list_item_left}>
                            <div className={styles.list_item_icon}>{CalendarIcon()}</div>
                            <p className={styles.list_item_label}>Datum</p>
                        </div>

                        <p className={styles.list_item_value}>10. Dezember 2025</p>
                    </li>
                )}

                {props.activeConfig.showTime && (
                    <li className={styles.list_item}>
                        <div className={styles.list_item_left}>
                            <div className={styles.list_item_icon}>{ClockIcon()}</div>
                            <p className={styles.list_item_label}>Uhrzeit</p>
                        </div>

                        <p className={styles.list_item_value}>10:00 Uhr</p>
                    </li>
                )}

                {props.activeConfig.showName && (
                    <li className={styles.list_item}>
                        <div className={styles.list_item_left}>
                            <div className={styles.list_item_icon}>{UserIcon()}</div>
                            <p className={styles.list_item_label}>Name</p>
                        </div>

                        <p className={styles.list_item_value}>Max Musterman</p>
                    </li>
                )}

                {props.activeConfig.showPhoneNumer && (
                    <li className={styles.list_item}>
                        <div className={styles.list_item_left}>
                            <div className={styles.list_item_icon}>{PhoneIcon()}</div>
                            <p className={styles.list_item_label}>Telefon</p>
                        </div>

                        <p className={styles.list_item_value}>+49 176 12345678</p>
                    </li>
                )}
            </ul>
        </div>
    );
};