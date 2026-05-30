using System;
using System.Net.Http;
using System.Text.Json;
using System.Threading.Tasks;

public class AttendanceResponse
{
    public string Status { get; set; }
    public Attendance[] Data { get; set; }
}

public class Attendance
{
    public int PersonId { get; set; }
    public string PersonName { get; set; }
    public int Status { get; set; }
    public long Timestamp { get; set; }
    public string Tanggal { get; set; }
    public string JamMasuk { get; set; }
    public string JamPulang { get; set; }
    public int ShiftId { get; set; }
    public string ShiftName { get; set; }
    public string ShiftSchedule { get; set; }
    public int TerlambatMenit { get; set; }
    public int LemburMenit { get; set; }
    public string StatusText { get; set; }
}

class Program
{
    static async Task Main()
    {
        var url = "https://your.api.com/attendance";

        using var client = new HttpClient();

        client.DefaultRequestHeaders.Add("X-API-Key", "YOUR_API_KEY");

        var response = await client.PostAsync(url, null);

        response.EnsureSuccessStatusCode();

        var json = await response.Content.ReadAsStringAsync();

        var result = JsonSerializer.Deserialize<AttendanceResponse>(
            json,
            new JsonSerializerOptions
            {
                PropertyNameCaseInsensitive = true
            });

        Console.WriteLine($"Status: {result?.Status}");

        if (result?.Data != null)
        {
            foreach (var item in result.Data)
            {
                Console.WriteLine($"Name: {item.PersonName}");
                Console.WriteLine($"Date: {item.Tanggal}");
                Console.WriteLine($"Check In: {item.JamMasuk}");
                Console.WriteLine($"Check Out: {item.JamPulang}");
                Console.WriteLine($"Status: {item.StatusText}");
                Console.WriteLine("---------------------");
            }
        }
    }
}