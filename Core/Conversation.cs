using MimeKit;

namespace Core;

public class Conversation(MimeMessage message)
{
    private static string InternetAddressListToString(InternetAddressList address)
    {
        return address.Count == 0 ? "Unknown" : address[0].Name;
    }

    public readonly string From = InternetAddressListToString(message.From);
    public string Subject => message.Subject;
    public string Content => message.TextBody;
    public DateTime Date => message.Date.LocalDateTime;

    public override string ToString()
    {
        return $"From {From} at {Date}:\n{Subject}";
    }
}