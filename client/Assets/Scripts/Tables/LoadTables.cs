
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
public class LoadTables
{
    /// <summary>
    /// Init ALL tables
    /// </summary>
	public static void Init()
    {
        FirstTableManager.Instance.InitTable();
        SecondTableManager.Instance.InitTable();
    }
}
