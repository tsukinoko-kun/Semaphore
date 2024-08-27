import "./inbox.css"
import Await from "./Await"
import { GetConversations } from "../wailsjs/go/main/App"

export default function Inbox() {
    return (
        <div id="inbox" className="h-full">
            <ConversationsList />
        </div>
    )
}

function ConversationsList() {
    return (
        <Await
            promise={GetConversations}
            then={(props) => (
                <ol
                    id="conversation_list"
                    className="m-4 overflow-y-auto overflow-x-clip rounded-md border border-solid border-neutral-700 bg-neutral-995"
                >
                    {props.value.map((conv) => (
                        <li key={conv.subject} className="cursor-pointer p-4 grayscale hover:bg-neutral-900">
                            <p className="text-balance text-white">{conv.subject}</p>
                            <p
                                className="max-w-full overflow-hidden overflow-ellipsis whitespace-nowrap text-neutral-400"
                                dangerouslySetInnerHTML={{ __html: conv.lastMessage }}
                            />
                        </li>
                    ))}
                </ol>
            )}
        />
    )
}
