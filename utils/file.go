package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func ChangeExcelFile(orgPath, destPath string) error {
	// 加载 .xls 文件
	xlsFile, err := excelize.OpenFile(orgPath)
	if err != nil {
		return err
	}

	// 创建新的 .xlsx 文件
	xlsxFile := excelize.NewFile()

	// 复制工作表到新文件
	for _, sheetName := range xlsFile.GetSheetMap() {
		rows := xlsFile.GetRows(sheetName)

		// 将原始数据写入新文件
		for rowIndex, row := range rows {
			for columnIndex, cell := range row {
				xlsxFile.SetCellValue(sheetName, excelize.ToAlphaString(columnIndex+1)+strconv.Itoa(rowIndex+1), cell)
			}
		}

		// 设置每个工作表的默认列宽（可选）
		xlsxFile.SetColWidth(sheetName, "A", "ZZ", 15)
	}

	// 保存新的 .xlsx 文件
	err = xlsxFile.SaveAs(destPath)
	if err != nil {
		return err
	}

	return nil
}
