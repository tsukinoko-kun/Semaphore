import { type FC, JSX, useEffect, useState } from "react"

export type ErrorProps = {
    err: Error
}

type Props<T> = {
    // a function that returns the promise to be awaited
    promise: () => Promise<T>
    // a function that returns the JSX element to be rendered when the promise is rejected
    catch?: FC<ErrorProps>
    // a function that returns the JSX element to be rendered when the promise is resolved
    then: FC<{ value: T }>
    // a function that returns the JSX element to be rendered while the promise is pending
    pending?: FC
}

type State = "pending" | "resolved" | "rejected"

export default function Await<T>(props: Props<T>): JSX.Element {
    const [promise, setPromise] = useState<Promise<T> | null>(null)
    const [state, setState] = useState<State>("pending")
    const [result, setResult] = useState<T | null>(null)
    const [error, setError] = useState<Error | null>(null)

    const Then = props.then
    const Catch = props.catch || ((props: ErrorProps) => <p className="text-red-500">{props.err.message}</p>)
    const Pending = props.pending || (() => <div>Loading...</div>)

    useEffect(() => {
        if (promise !== null) {
            return
        }
        console.log("Await: useEffect")
        const p = props.promise()
        setPromise(p)
        p.then((res) => {
            setResult(res)
            setState("resolved")
        }).catch((err) => {
            setError(err)
            setState("rejected")
        })
    }, [])

    switch (state) {
        case "pending":
            return <Pending />
        case "resolved":
            return <Then value={result as T} />
        case "rejected":
            return <Catch err={error as Error} />
    }
}
