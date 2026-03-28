import styles from "./style.module.scss";
import {Divider} from "../divider";

interface IProps {
    name: string;
    callbackAction?: () => void;
}

export const ViewHeader = (props: IProps) => {
    return (
        <div className={styles.header}>
            <div className={styles.header_row}>
                {props.callbackAction && (
                    <button onClick={props.callbackAction}>
                        c-
                    </button>
                )}


                <h1>{props.name}</h1>
            </div>

            <Divider/>
        </div>
    )
}