from concurrent.futures.process import ProcessPoolExecutor


class MProcessPool:
    def __init__(self):
        self.db = object()

    def pull(self):
        pass

    def run(self):
        with ProcessPoolExecutor(max_workers=4) as pool:
            for process_num, table in enumerate(self.db.table):
                pool.submit(self.pull, process_num, table)