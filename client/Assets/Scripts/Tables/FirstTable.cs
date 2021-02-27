using System.Collections.Generic;
using System.IO;
using System.Text;
using UnityEngine;

public class FirstTable
{
    public int ID;
    public string F_Name;
    public int Age;
    public string S_Name;
    public FirstTable(string line)
    {
        string[] fileds = line.Split('	');
        ID = int.Parse(fileds[0]);
        F_Name = fileds[1];
        Age = int.Parse(fileds[2]);
        S_Name = fileds[3];
     }
}

public class FirstTableManager
{
    Dictionary<int, FirstTable> data = new Dictionary<int, FirstTable>();

    public void InitTable()
    {
        string tableDir = PathUtil.GetTableDataPath();
        string tableName = "FirstTable";
        string tableDataPath = string.Format("{0}{1}.txt", tableDir, tableName);
        StreamReader sr = new StreamReader(tableDataPath, Encoding.UTF8);
        string line;
        while ((line = sr.ReadLine()) != null)
        {
            line = line.Trim();
            if (line.Length > 0)
            {
                FirstTable rowData = new FirstTable(line);
                data.Add(rowData.ID, rowData);
            }
            else
            {
                continue;
            }
        }
    }

    public FirstTable GetDataByID(int id)
    {
        FirstTable rowData = null;
        data.TryGetValue(id, out rowData);
        return rowData;
    }

    public FirstTable GetDataByName(string name)
    {
        foreach(KeyValuePair<int, FirstTable> kp in data)
        {
            if (kp.Value.Name == name) return kp.Value;
        }
        return null;
    }
    private FirstTableManager() { }
    public static readonly FirstTableManager Instance = new FirstTableManager();
} 