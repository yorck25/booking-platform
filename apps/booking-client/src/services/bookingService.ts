import { Config } from "../lib/config";
import { NetworkAdapter, NewDefaultHeader } from "../lib/networkAdapter";
import {type IReservationSlot} from "../models/barber";
import { useSearchParams } from "react-router-dom";

class CBookingService {
    private BarberId = "";
    private BASE_URL = Config.BookingServiceBaseUrl;
    private readonly BARBER_ID_STORAGE_KEY = "barberId";

    setBarberId(id: string) {
        this.BarberId = id;
        sessionStorage.setItem(this.BARBER_ID_STORAGE_KEY, id);
    }

    getBarberId(): string {
        if (this.BarberId) {
            return this.BarberId;
        }

        const idFromUrl = this.getBarberIdFromUrl();

        if (idFromUrl) {
            this.setBarberId(idFromUrl);
            return idFromUrl;
        }

        const idFromSession = sessionStorage.getItem(this.BARBER_ID_STORAGE_KEY);

        if (idFromSession) {
            this.BarberId = idFromSession;
            return idFromSession;
        }

        console.error("No BarberId provided");
        return "";
    }

    private getBarberIdFromUrl(): string {
        const searchParams = new URLSearchParams(window.location.search);
        const barberId = searchParams.get("barberId");

        return barberId || "";
    }

    // -----------------
    // Bookings
    // -----------------
    createBooking(): boolean {
        const headers = NewDefaultHeader();

        const requestOptions: RequestInit = {
            method: NetworkAdapter.POST,
            headers: headers
        }

        fetch(`${this.BASE_URL}/bookings`, requestOptions).then(r => {
            if(r.status !== 200) {
                console.warn(r.body);
                return false;
            }

            return r.status === 200;
        })

        return false;
    }

    cancelBooking(bookingId: string, customerPhoneNumber: string, cancelReason: string): boolean {
        const headers = NewDefaultHeader();

        const requestOptions: RequestInit = {
            method: NetworkAdapter.PUT,
            headers: headers,
            body: JSON.stringify{
                bookingId: bookingId,
                customerPhoneNumber: customerPhoneNumber,
                cancelReason: cancelReason,
            }
        }

        fetch(`${this.BASE_URL}/bookings/cancel`, requestOptions).then(r => {
            if(r.status !== 200) {
                console.warn(r.body);
                return false;
            }

            return r.status === 200;
        })

        return false;
    }

    // -----------------
    // Barber
    // -----------------

    loadAvailableService(barberId: string) {
        const headers = NewDefaultHeader();

        const requestOptions: RequestInit = {
            method: NetworkAdapter.GET,
            headers: headers
        }

        console.log(this.BASE_URL)

        fetch(`${this.BASE_URL}/barber/services/list?barberId=${barberId}`, requestOptions).then(r => {
            if(r.status !== 200) {
                console.warn(r.body);
                return false;
            }
            return r.json();
        }).then(data => {
            console.log(data)
        })

        return false;
    }

    loadBarberAvailability(barberId: string, serviceId: string, bookingDate: string): IReservationSlot[] {
        const headers = NewDefaultHeader();

        const requestOptions: RequestInit = {
            method: NetworkAdapter.GET,
            headers: headers
        }

        fetch(`${this.BASE_URL}/barber/availability?barberId=${barberId}&serviceId=${serviceId}&bookingDate=${bookingDate}`, requestOptions).then(r => {
            if(r.status !== 200) {
                console.warn(r.body);
                return false;
            }
            return r.json();
        }).then((data: IReservationSlot[]) => {
            return data;
        })

        return [];
    }
}

export const BookingService = new CBookingService();