import { FirstPage } from "../wailsjs/go/main/App"
import { useNavigate } from "react-router-dom"
import { useEffect } from "react"

export default function Init() {
    const navigate = useNavigate()
    useEffect(() => {
        FirstPage()
            .then((firstPage) => {
                navigate(firstPage)
            })
            .catch((err) => {
                alert(err)
            })
    })

    return null
}
