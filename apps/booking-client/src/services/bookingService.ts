import { Config } from "../lib/config";
import { NetworkAdapter, NewDefaultHeader } from "../lib/networkAdapter";

export class BookingService {
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
    // Bookings
    // -----------------


}