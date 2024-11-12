using System.Net.Security;
using System.Security.Cryptography.X509Certificates;
using MailKit;
using MailKit.Net.Imap;
using MailKit.Security;

namespace Core;

public class Conn : IDisposable
{
    private bool _disposed = false;
    private readonly IEnumerable<Account> _accounts;

    public Conn()
    {
        _accounts = [new Account("127.0.0.1", 1143, 1025, true, "hello@frankmayer.dev", "AoWS8-s6DT4lR1hrWPHG6w")];
    }

    public void Dispose()
    {
        if (!_disposed) return;
        _disposed = true;
        foreach (var account in _accounts)
        {
            account.Dispose();
        }
    }

    ~Conn()
    {
        Dispose();
    }

    public async IAsyncEnumerable<Conversation> InboxAsync()
    {
        foreach (var account in _accounts)
        {
            await foreach (var m in account.InboxAsync())
            {
                yield return m;
            }
        }
    }
}