import redis
from py.sql_and_nosql.tools import get_project_settings


class SSDBClientHandler:
    def __init__(self):
        setting = get_project_settings()
        SSDB_PARAMS = setting.SSDB_PARAMS
        HOST = SSDB_PARAMS['host']
        PORT = SSDB_PARAMS['port']
        PARAMS = SSDB_PARAMS['password']

        self.client = redis.Redis(host=HOST, port=PORT, password=PARAMS,
                             decode_responses=True)

    def execute_command(self, command, *args):
        return self.client.execute_command(command, *args)