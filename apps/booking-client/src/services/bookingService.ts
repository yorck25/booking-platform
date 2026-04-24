import { Config } from "../lib/config";
import { NetworkAdapter, NewDefaultHeader } from "../lib/networkAdapter";
import {type IReservationSlot} from "../models/barber";

class CBookingService {
    public BarberId = "11111111-1111-1111-1111-111111111111";
    private BASE_URL = Config.BookingServiceBaseUrl;

    setBarberId(id: string) {
        this.BarberId = id;
    }

    getBarberId(): string {
        return this.BarberId;
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