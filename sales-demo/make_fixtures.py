"""Generate the demo input workbooks and product mapping, deterministically.

Run: python make_fixtures.py   (writes into ./data)
The generated files are committed; this script is how they were made and
documents each deliberate defect.
"""
import csv
import datetime as dt
from pathlib import Path

from openpyxl import Workbook

HEADER = ["Order ID", "Order Date", "Customer", "Product", "Category",
          "Quantity", "Unit Price", "Amount", "Status"]

# (row comment, values). Excel row = index + 2 (row 1 is the header).
SEPT = [
    ("clean, ISO date",            ["A-1001", "2026-09-01", "Acme", "Widget Pro", "Hardware", 2, "10.50", "21.00", "Paid"]),
    ("clean, DD/MM/YYYY date",     ["A-1002", "02/09/2026", "Globex", "widget pro", "Hardware", 1, "10.50", "10.50", "Paid"]),
    ("clean, real Excel date",     ["A-1003", dt.datetime(2026, 9, 3), "Initech", "Gadget-X", "Hardware", 3, "4.10", "12.30", "Open"]),
    ("alias product name",         ["A-1004", "2026-09-04", "Acme", "GADGET X", "Hardware", 1, "4.10", "4.10", "Paid"]),
    ("identical duplicate (1/2)",  ["A-1005", "2026-09-05", "Hooli", "Support Plan", "Service", 1, "99.99", "99.99", "Paid"]),
    ("identical duplicate (2/2)",  ["A-1005", "2026-09-05", "Hooli", "Support Plan", "Service", 1, "99.99", "99.99", "Paid"]),
    ("conflicting duplicate (1/2)",["A-1006", "2026-09-06", "Umbrella", "Widget Pro", "Hardware", 2, "10.50", "21.00", "Paid"]),
    ("conflicting duplicate (2/2)",["A-1006", "2026-09-06", "Umbrella", "Widget Pro", "Hardware", 3, "10.50", "31.50", "Paid"]),
    ("invalid date",               ["A-1007", "31/02/2026", "Acme", "Widget Pro", "Hardware", 1, "10.50", "10.50", "Paid"]),
    ("missing customer",           ["A-1008", "2026-09-08", None, "Widget Pro", "Hardware", 1, "10.50", "10.50", "Paid"]),
    ("zero quantity",              ["A-1009", "2026-09-09", "Acme", "Widget Pro", "Hardware", 0, "10.50", "0.00", "Paid"]),
    ("fractional quantity",        ["A-1010", "2026-09-10", "Acme", "Widget Pro", "Hardware", 1.5, "10.50", "15.75", "Paid"]),
    ("negative unit price",        ["A-1011", "2026-09-11", "Acme", "Widget Pro", "Hardware", 1, "-1.00", "-1.00", "Paid"]),
    ("amount mismatch",            ["A-1012", "2026-09-12", "Acme", "Widget Pro", "Hardware", 2, "10.50", "20.00", "Paid"]),
    ("non-numeric price",          ["A-1013", "2026-09-13", "Acme", "Widget Pro", "Hardware", 1, "ten", "10.00", "Paid"]),
    ("boundary: zero price",       ["A-1014", "2026-09-14", "Acme", "Free Sample", "Hardware", 1, "0.00", "0.00", "Paid"]),
    ("decimal trap 0.1 x 3",       ["A-1015", "2026-09-15", "Acme", "Gadget-X", "Hardware", 3, "0.10", "0.30", "Paid"]),
    ("unmapped product",           ["A-1016", "2026-09-16", "Acme", "Mystery Box", "Hardware", 1, "5.00", "5.00", "Paid"]),
]

OCT = [
    ("", ["B-2001", "2026-10-01", "Acme", "Widget Pro", "Hardware", 4, "10.50", "42.00", "Paid"]),
    ("", ["B-2002", "03/10/2026", "Globex", "support plan", "Service", 2, "99.99", "199.98", "Open"]),
    ("", ["B-2003", "2026-10-04", "Initech", "Gadget X", "Hardware", 5, "4.10", "20.50", "Paid"]),
]

MAPPING = [
    ("raw_name", "canonical_name"),
    ("Widget Pro", "Widget Pro"), ("widget pro", "Widget Pro"),
    ("Gadget-X", "Gadget X"), ("GADGET X", "Gadget X"), ("Gadget X", "Gadget X"),
    ("Support Plan", "Support Plan"), ("support plan", "Support Plan"),
    ("Free Sample", "Free Sample"),
]


def write_workbook(path, rows):
    wb = Workbook()
    ws = wb.active
    ws.title = "Orders"
    ws.append(HEADER)
    for _, values in rows:
        ws.append(values)
    wb.save(path)


def main():
    out = Path(__file__).parent / "data"
    out.mkdir(exist_ok=True)
    write_workbook(out / "sales_2026_09.xlsx", SEPT)
    write_workbook(out / "sales_2026_10.xlsx", OCT)
    with open(out / "product_mapping.csv", "w", newline="") as f:
        csv.writer(f, lineterminator="\n").writerows(MAPPING)


if __name__ == "__main__":
    main()
