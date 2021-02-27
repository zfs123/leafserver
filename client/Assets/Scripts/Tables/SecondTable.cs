using System.Collections.Generic;
using System.IO;
using System.Text;
using UnityEngine;

public class SecondTable
{
    public int ID;
    public string F_Name;
    public int Age;
    public string S_Name;
    public SecondTable(string line)
    {
        string[] fileds = line.Split('	');
        ID = int.Parse(fileds[0]);
        F_Name = fileds[1];
        Age = int.Parse(fileds[2]);
        S_Name = fileds[3];
     }
}

public class SecondTableManager
{
    Dictionary<int, SecondTable> data = new Dictionary<int, SecondTable>();

    public void InitTable()
    {
        string tableDir = PathUtil.GetTableDataPath();
        string tableName = "SecondTable";
        string tableDataPath = string.Format("{0}{1}.txt", tableDir, tableName);
        StreamReader sr = new StreamReader(tableDataPath, Encoding.UTF8);
        string line;
        while ((line = sr.ReadLine()) != null)
        {
            line = line.Trim();
            if (line.Length > 0)
            {
                SecondTable rowData = new SecondTable(line);
                data.Add(rowData.ID, rowData);
            }
            else
            {
                continue;
            }
        }
    }

    public SecondTable GetDataByID(int id)
    {
        SecondTable rowData = null;
        data.TryGetValue(id, out rowData);
        return rowData;
    }

    public SecondTable GetDataByName(string name)
    {
        foreach(KeyValuePair<int, SecondTable> kp in data)
        {
            if (kp.Value.Name == name) return kp.Value;
        }
        return null;
    }
    private SecondTableManager() { }
    public static readonly SecondTableManager Instance = new SecondTableManager();
} 