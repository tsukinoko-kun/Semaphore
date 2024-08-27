import {useState} from "react";
import {AddAccount} from "../wailsjs/go/main/App";
import {useNavigate} from "react-router-dom";
import {Eye, EyeSlash} from "react-bootstrap-icons";

export default function Login() {
    const navigate = useNavigate()
    const [displayName, setDisplayName] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [server, setServer] = useState("")
    const [port, setPort] = useState("")

    const [pwdVis, setPwdVis] = useState(false)

    const [error, setError] = useState("")

    return <div className="h-screen flex flex-col justify-center w-full max-w-xl m-auto">
        <h1 className="w-fit select-none">Login</h1>
        <form className="flex flex-col gap-8" onSubmit={(ev) => {
            ev.preventDefault()
            AddAccount(displayName, email, password, server, port).then((err: string) => {
                if (err) {
                    setError(err)
                    return
                }
                navigate("/inbox")
            }).catch((err: any) => {
                if (err instanceof Error) {
                    setError(err.message)
                } else if (typeof err === "string") {
                    setError(err)
                } else {
                    setError(JSON.stringify(err))
                }
            })
        }}>
            <label className="input">
                Display Name
                <input
                    type="text"
                    placeholder="John Doe"
                    value={displayName}
                    required
                    onChange={(ev) => {
                        setDisplayName(ev.target.value)
                    }}
                />
            </label>
            <label className="input">
                E-Mail Address
                <input
                    type="email"
                    placeholder="hello@example.com"
                    value={email}
                    required
                    onChange={(ev) => {
                        setEmail(ev.target.value)
                    }}
                />
            </label>
            <label className="input multi">
                <p>Password</p>
                <input
                    type={pwdVis ? "text" : "password"}
                    value={password}
                    required
                    onChange={(ev) => {
                        setPassword(ev.target.value)
                    }}
                />
                <span onClick={() => setPwdVis(!pwdVis)} className="cursor-pointer">
                {pwdVis ? <EyeSlash/> : <Eye/>}
                </span>
            </label>
            <label className="input multi">
                <p>Server</p>
                <input
                    type="text"
                    placeholder="imap.example.com"
                    autoComplete="off"
                    autoCapitalize={"none"}
                    autoCorrect={"off"}
                    value={server}
                    required
                    onChange={(ev) => {
                        setServer(ev.target.value)
                    }}
                />
                <span>:</span>
                <input
                    type="number"
                    autoComplete="off"
                    autoCapitalize={"none"}
                    autoCorrect={"off"}
                    placeholder="12345"
                    value={port}
                    required
                    onChange={(ev) => {
                        setPort(ev.target.value)
                    }}
                    min={1}
                    max={65535}
                />
            </label>
            <button className="btn mt-8" type="submit">Login</button>
        </form>
        {error && <p className="text-red-500 whitespace-pre my-8 select-text cursor-text">{error}</p>}
    </div>;
}