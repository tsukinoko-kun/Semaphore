namespace Core;

public interface IPasswd
{
    public void SetPassword(string accountName, string password);
    public string? GetPassword(string accountName);
}