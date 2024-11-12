using System;
using System.Runtime.InteropServices;

namespace Core
{
    public class Keychain : IPasswd
    {
        [DllImport("/System/Library/Frameworks/Security.framework/Security")]
        private static extern int SecKeychainFindGenericPassword(
            IntPtr keychainOrArray,
            int serviceNameLength,
            string serviceName,
            int accountNameLength,
            string accountName,
            out uint passwordLength,
            out IntPtr passwordData,
            IntPtr itemRef
        );

        [DllImport("/System/Library/Frameworks/Security.framework/Security")]
        private static extern int SecKeychainAddGenericPassword(
            IntPtr keychainOrArray,
            int serviceNameLength,
            string serviceName,
            int accountNameLength,
            string accountName,
            int passwordLength,
            string passwordData,
            out IntPtr itemRef
        );

        private const string ServiceName = "dev.frankmayer.semaphore";

        public void SetPassword(string accountName, string password)
        {
            int result = SecKeychainAddGenericPassword(
                IntPtr.Zero,                    // Use the default Keychain
                ServiceName.Length,             // Length of the service name
                ServiceName,                    // Service name
                accountName.Length,             // Length of the account name
                accountName,                    // Account name
                password.Length,                // Length of the password
                password,                       // Password data
                out _                           // Item reference
            );

            if (result != 0)
            {
                throw new Exception($"Failed to set password. Error code: {result}");
            }
        }

        public string? GetPassword(string accountName)
        {
            uint passwordLength = 0;
            IntPtr passwordData = IntPtr.Zero;

            int result = SecKeychainFindGenericPassword(
                IntPtr.Zero,
                ServiceName.Length,
                ServiceName,
                accountName.Length,
                accountName,
                out passwordLength,
                out passwordData,
                IntPtr.Zero
            );

            if (result == 0 && passwordLength > 0)
            {
                byte[] passwordBytes = new byte[passwordLength];
                Marshal.Copy(passwordData, passwordBytes, 0, (int)passwordLength);

                // Release unmanaged memory allocated by SecKeychainFindGenericPassword
                Marshal.FreeCoTaskMem(passwordData);

                return System.Text.Encoding.UTF8.GetString(passwordBytes);
            }

            return null;
        }
    }
}
