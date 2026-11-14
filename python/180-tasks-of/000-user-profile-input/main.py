class Profile:
    def __init__(self, name: str, age: int, department: str) -> None:
        self.name = name
        self.age = age
        self.department = department

    def __str__(self) -> str:
        return f"Profile: (\n\tname: {self.name}\n\tage: {self.age}\n\tdepartment: {self.department}\n)"


class ProfileParser:
    def parse(self) -> Profile | None:
        min_name_len = 3
        min_age = 0
        max_age = 120
        min_department_len = 2
        allowed_departments = ("finance", "management", "it")

        name = input("input name:\n").strip()
        if len(name) < min_name_len:
            print(f"name should be equal or longer than {min_name_len} chars")
            return

        age = 0
        try:
            age = int(input("input age:\n"))
        except ValueError:
            print(f"age should a be number between {min_age} and {max_age}")
            return

        if age < min_age or age > max_age:
            print(f"age should a be number between {min_age} and {max_age}")
            return

        department = input("input department:\n")
        if len(department) < min_department_len:
            print(f"department should be equal or longer than {min_department_len} chars")
            return

        if department not in allowed_departments:
            print(f"department should be one of {allowed_departments}")
            return

        return Profile(name, age, department)


def main():
    parser = ProfileParser()
    profile = None
    while profile is None:
        profile = parser.parse()
    print(profile)



if __name__ == "__main__":
    main()
