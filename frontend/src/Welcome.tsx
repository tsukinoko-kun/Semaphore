import { Link } from "react-router-dom"
import { Envelope } from "react-bootstrap-icons"

export default function Welcome() {
    return (
        <div className="m-auto flex h-screen w-fit flex-col justify-center">
            <h1 className="w-fit">Semaphore</h1>
            <p className="w-fit">not just an email client</p>
            <p className="w-fit">your messenger for the professional world</p>
            <Link draggable={false} to={"/login"} className="btn mt-8">
                <Envelope className="mr-4 inline" />
                Login
            </Link>
        </div>
    )
}
