import './app.css'
import {ServicePage} from "./routes/servicePage";
import {BrowserRouter, Routes, Route, useNavigate} from "react-router-dom";
import {ReservationPage} from "./routes/reservationPage";
import {PersonalDataPage} from "./routes/personalDataPage";
import {useEffect} from "preact/hooks";
import {SuccessBookingView} from "./routes/successBookingView";

export function App() {
    return (
        <>
            <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<RedirectToServices />} />                        <Route path="/services" element={<ServicePage />} />
                        <Route path="/reservation" element={<ReservationPage />} />
                        <Route path="/personal-data" element={<PersonalDataPage />} />
                        <Route path="/success-booking" element={<SuccessBookingView />} />
                    </Routes>
            </BrowserRouter>
        </>
    );
}

const RedirectToServices = () => {
    const navigate = useNavigate();

    useEffect(() => {
        navigate("/services");
    }, []);

    return (
        <>
            <p>Redirecting to services...</p>
            <a href="/services">Click here if not redirected</a>
        </>
    );
};