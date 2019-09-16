def myreadlines(f, newline):
    buf = ''
    while True:
        while newline in buf:
            pos = buf.index(newline)
            yield buf[:pos]
            buf = buf[pos + len(newline):]
        chunk = f.read(4096)

        if not chunk:
            yield buf
            break
        buf += chunk
        

if __name__ == '__main__':
    # 文件中的分隔符
    flite = r"\n"
    with open("contain.txt") as f:
        for line in myreadlines(f, flite):
            print(line)