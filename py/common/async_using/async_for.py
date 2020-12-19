import asyncio


class AsyncIteratorWrapper:
    def __init__(self, obj):
        self._it = iter(obj)

    def __aiter__(self):
        return self

    async def __anext__(self):
        try:
            value = next(self._it)
        except StopIteration:
            raise StopAsyncIteration
        return value


async def run_for(string):
    async for letter in AsyncIteratorWrapper(string):
        print(letter)


def main():
    loop = asyncio.get_event_loop()
    tasks = [asyncio.ensure_future(each) for each in [run_for("abcd"), ]]
    loop.run_until_complete(asyncio.wait(tasks))


if __name__ == '__main__':
    main()