using System.Net.Security;
using System.Security.Cryptography.X509Certificates;
using MailKit;
using MailKit.Net.Imap;
using MailKit.Net.Smtp;
using MailKit.Security;

namespace Core;

public class Account : IDisposable
{
    private readonly ImapClient _imap;
    private readonly SmtpClient _smtp;
    private bool _disposed;

    public Account(string host, int imapPort, int smtpPort, bool smtpSsl, string username, string password)
    {
        _smtp = new SmtpClient();
        _smtp.ServerCertificateValidationCallback = ServerCertificateValidationCallback;
        _smtp.Timeout = 5000;
        _smtp.Connect(host, smtpPort, smtpSsl ? SecureSocketOptions.SslOnConnect : SecureSocketOptions.Auto);
        _smtp.Authenticate(username, password);

        _imap = new ImapClient();
        try
        {
            _imap.ServerCertificateValidationCallback = ServerCertificateValidationCallback;
            _imap.Timeout = 5000;
            _imap.Connect(host, imapPort, SecureSocketOptions.StartTlsWhenAvailable);
            _imap.Authenticate(username, password);
        }
        catch (Exception ex)
        {
            if (_smtp.IsConnected)
            {
                _smtp.Disconnect(true);
            }

            if (_imap.IsConnected)
            {
                _imap.Disconnect(true);
            }

            throw ex;
        }
    }

    public void Dispose()
    {
        if (_disposed) return;
        _disposed = true;
        _imap.Dispose();
        _smtp.Dispose();
    }

    ~Account()
    {
        Dispose();
    }

    private static bool ServerCertificateValidationCallback(object sender, X509Certificate? certificate,
        X509Chain? chain, SslPolicyErrors sslPolicyErrors)
    {
        return true;
    }

    public async IAsyncEnumerable<Conversation> InboxAsync()
    {
        var inbox = _imap.Inbox;
        if (inbox is null) yield break;

        await inbox.OpenAsync(FolderAccess.ReadOnly);

        for (var i = 0; i < inbox.Count; i++)
        {
            var m = await inbox.GetMessageAsync(i);
            if (m is null) continue;

            yield return new Conversation(m);
        }

        await inbox.CloseAsync();
    }
}