import csv
import random
from datetime import date, timedelta

output_file = "txns.csv"

fields = ["Id", "Date", "Amount"]
start_date = date(2023, 1, 1)
end_date = date(2023, 10, 31)
num_rows = 100


def generate_random_date():
    random_datetime = start_date + timedelta(days=random.randint(0, (end_date - start_date).days))
    return random_datetime.strftime("%Y-%m-%d")


def generate_random_amount():
    value = round(random.uniform(1, 100), 2)
    return f"+{value}" if random.choice([True, False]) else f"-{value}"


with open(output_file, mode="w", newline="") as csv_file:
    writer = csv.DictWriter(csv_file, fieldnames=fields)

    writer.writeheader()
    for i in range(1, num_rows + 1):
        random_date = generate_random_date()
        random_amount = generate_random_amount()
        writer.writerow({"Id": i, "Date": random_date, "Amount": random_amount})

print(output_file)
