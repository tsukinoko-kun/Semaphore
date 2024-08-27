import {Link} from "react-router-dom";
import {Envelope} from "react-bootstrap-icons";

export default function Welcome() {
    return <div className="h-screen flex flex-col justify-center w-fit m-auto">
        <h1 className="w-fit">Semaphore</h1>
        <p className="w-fit">not just an email client</p>
        <p className="w-fit">your messenger for the professional world</p>
        <Link draggable={false} to={"/login"} className="btn mt-8">
            <Envelope className="inline mr-4"/>
            Login
        </Link>

    </div>;
}