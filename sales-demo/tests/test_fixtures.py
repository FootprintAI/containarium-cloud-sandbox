"""Baseline: the committed inputs exist and have the documented shape."""
import csv
from pathlib import Path

from openpyxl import load_workbook

DATA = Path(__file__).resolve().parent.parent / "data"
COLUMNS = ["Order ID", "Order Date", "Customer", "Product", "Category",
           "Quantity", "Unit Price", "Amount", "Status"]


def test_september_workbook_shape():
    ws = load_workbook(DATA / "sales_2026_09.xlsx", read_only=True)["Orders"]
    rows = list(ws.iter_rows(values_only=True))
    assert list(rows[0]) == COLUMNS
    assert len(rows) - 1 == 18


def test_mapping_has_header_and_rows():
    with open(DATA / "product_mapping.csv", newline="") as f:
        rows = list(csv.reader(f))
    assert rows[0] == ["raw_name", "canonical_name"]
    assert len(rows) > 1
