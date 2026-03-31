import { Config } from "../lib/config";
import { NetworkAdapter, NewDefaultHeader } from "../lib/networkAdapter";

class CBookingService {
    public BarberId = "11111111-1111-1111-1111-111111111111";
    private BASE_URL = Config.BookingServiceBaseUrl;

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
}

export const BookingService = new CBookingService();