import importlib


def get_project_settings():
    modele = importlib.import_module("setting")
    return modele


if __name__ == '__main__':
    get_project_settings()