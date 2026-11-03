# sales-demo

Seed for the E2E-EXCEL-001 demo: a coding agent builds `sales_processor`, a
CLI that turns monthly sales workbooks into validated, standardized Excel
outputs. This directory holds only the scaffold: input fixtures, a baseline
test, CI, and agent guidance. The tool itself is what the agent builds.

`make_fixtures.py` documents each deliberate defect in `data/sales_2026_09.xlsx`
(18 order rows: mixed date formats, an identical duplicate, a conflicting
duplicate, an invalid date, a missing field, bad quantities and prices, an
amount mismatch, a 0.1 x 3 decimal trap, an unmapped product).
`data/sales_2026_10.xlsx` is a second, small month for the "no code changes"
check.

```
cd sales-demo
pip install -r requirements.txt
python -m pytest -q
```
