import styles from "./style.module.scss";
import { ViewHeader } from "../../components/viewHeader";
import { dummyServices, type IService, type ServiceCategory } from "../../models/service.ts";
import { useMemo, useState } from "preact/hooks";
import {ChevronIcon, ClockIcon} from "../../components/icons";
import { Divider } from "../../components/divider";
import {BookingService} from "../../services/bookingService";
import {useEffect} from "preact/hooks"

export const ServicePage = () => {
    const [selectedServiceId, setSelectedServiceId] = useState<string>("22222222-2222-2222-2222-222222222222");
    const [note, setNote] = useState<string>("");

    useEffect(() => {
        BookingService.loadAvailableService("11111111-1111-1111-1111-111111111111");
    }, []);

    const groupedServices = useMemo(() => {
        const sortServices = (category: ServiceCategory) =>
            dummyServices
                .filter((service) => service.category === category && service.active && !service.deleted)
                .sort((a, b) => a.sortOrder - b.sortOrder);

        return {
            female: sortServices("female"),
            male: sortServices("male"),
            other: sortServices("other"),
        };
    }, []);

    const renderSection = (title: string, services: IService[]) => {
        if (!services.length) {
            return null;
        }

        return (
            <section className={styles.section}>
                <h2 className={styles.section_title}>{title}</h2>

                <ul className={styles.service_list}>
                    {services.map((service) => {
                        const isSelected = selectedServiceId === service.id;

                        return (
                            <li key={service.id}>
                                <button
                                    type="button"
                                    className={`${styles.service_card} ${isSelected ? styles.service_card_selected : ""}`}
                                    onClick={() => setSelectedServiceId(service.id)}
                                >
                                    <div className={styles.service_card_top}>
                                        <p className={styles.service_name}>{service.displayName}</p>

                                        {isSelected ? (
                                            <span className={styles.selected_badge}>Ausgewählt</span>
                                        ) : (
                                            <span className={styles.select_action}>
                                                Auswählen
                                                <div className={styles.chevron}>{ChevronIcon()}</div>
                                            </span>
                                        )}
                                    </div>

                                    <div className={styles.service_meta}>
                                        <span className={styles.service_meta_icon}>{ClockIcon()}</span>
                                        <span>{service.durationMinutes} Min.</span>
                                    </div>
                                </button>
                            </li>
                        );
                    })}
                </ul>
            </section>
        );
    };

    return (
        <div className={styles.service_view}>
            <ViewHeader name={"Service Auswählen"} />

            <div className={styles.content}>
                <p className={styles.subtitle}>
                    Wählen Sie ihre gewünschte Behandlung aus oder hinterlassen Sie eine Notiz für besondere Wünsche
                </p>

                <div className={styles.render_service_section}>
                    {renderSection("Damen", groupedServices.female)}
                    {renderSection("Herren", groupedServices.male)}
                    {renderSection("Sonstiges", groupedServices.other)}
                </div>

                <Divider />

                <div className={styles.note_section}>
                    <h2 className={styles.section_title}>Notiz Hinzufügen</h2>

                    <textarea
                        className={styles.note_input}
                        placeholder="Zum Beispiel: Nur Spitzen schneiden....."
                        value={note}
                        onInput={(e) => setNote((e.target as HTMLTextAreaElement).value)}
                    />
                </div>
            </div>

            <div className={styles.footer}>
                <button type="button" className={styles.button}>
                    Weiter
                </button>
            </div>
        </div>
    );
};