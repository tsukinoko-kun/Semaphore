using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Threading.Tasks;
using Avalonia.Controls;
using Core;
using Microsoft.Extensions.DependencyInjection;

namespace Desktop;

public partial class MainWindow : Window
{
    private readonly Conn _conn = Program.ServiceProvider.GetRequiredService<Conn>();
    private readonly ObservableCollection<string> _subjects = [];
    private readonly List<Conversation> _conversations = [];

    public MainWindow()
    {
        InitializeComponent();
        MessagesListBox.ItemsSource = _subjects;

        _ = LoadMessagesAsync();
    }

    private async Task LoadMessagesAsync()
    {
        _subjects.Clear();
        _conversations.Clear();
        await foreach (var c in _conn.InboxAsync())
        {
            var inserted = false;
            for (var i = 0; i < _subjects.Count; i++)
            {
                if (_conversations[i].Date < c.Date)
                {
                    _conversations.Insert(i, c);
                    _subjects.Insert(i, c.ToString());
                    inserted = true;
                    break;
                }
            }

            if (!inserted)
            {
                _conversations.Add(c);
                _subjects.Add(c.ToString());
            }
        }
    }
}