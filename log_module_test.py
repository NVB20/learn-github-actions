import unittest
import os
import json
from main import create_log_file, get_log_number

class TestCreateLogFile(unittest.TestCase):
    def setUp(self):
        self.log_filename = "test_logtimes.json"

    def tearDown(self):
        if os.path.exists(self.log_filename):
            os.remove(self.log_filename)

    def test_get_log_number(self):
        log_number = get_log_number(self.log_filename)
        self.assertEqual(log_number, 1)

        with open(self.log_filename, 'a') as log_file:
            log_file.write(json.dumps({"timestamp": "2023-08-21 12:34:56", "log_number": 2}) + "\n")

        log_number = get_log_number(self.log_filename)
        self.assertEqual(log_number, 2)

    def test_get_log_number_file_not_found(self):
        log_number = get_log_number("nonexistent_log.json")
        self.assertEqual(log_number, 1)

if __name__ == "__main__":
    unittest.main()