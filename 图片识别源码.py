
# 中文识别接口
from tkinter import *
import tkinter.filedialog
from PIL import ImageGrab
from time import sleep
from os import remove
from requests import post


def sb(path):
    multiple_files = {'pic': ('1111111.jpg', open(path, 'rb'), 'image/jpg')}
    resp = post(r'http://ocr.shouji.sogou.com/v2/ocr/json', files=multiple_files)
    str_json = resp.json()
    confirmLabel.delete('1.0', 'end')  # 清空文本框的内容
    for i in range(150):
        try:
            confirmLabel.insert(END, str_json['result'][i]['content'], )  # 打印到GUI界面
            window.update()
        except:
            pass


def xz():
    filenames = tkinter.filedialog.askopenfilenames()  # 选择上传的文件，
    sb(filenames[0])  # 获取到文件路径


def JieTu():

    class MyCapture:
        def __init__(self, png):
            # 变量X和Y用来记录鼠标左键按下的位置
            self.X = tkinter.IntVar(value=0)
            self.Y = tkinter.IntVar(value=0)
            # 屏幕尺寸
            screenWidth = window.winfo_screenwidth()
            screenHeight = window.winfo_screenheight()
            # 创建顶级组件容器
            self.top = tkinter.Toplevel(window, width=screenWidth, height=screenHeight)
            # 不显示最大化、最小化按钮
            self.top.overrideredirect(True)
            self.canvas = tkinter.Canvas(self.top, bg='white', width=screenWidth, height=screenHeight)
            # 显示全屏截图，在全屏截图上进行区域截图
            self.image = tkinter.PhotoImage(file=png)
            self.canvas.create_image(screenWidth // 2, screenHeight // 2, image=self.image)

            # 鼠标左键按下的位置
            def onLeftButtonDown(event):
                self.X.set(event.x)
                self.Y.set(event.y)
                # 开始截图
                self.sel = True

            self.canvas.bind('<Button-1>', onLeftButtonDown)

            # 鼠标左键移动，显示选取的区域
            def onLeftButtonMove(event):
                if not self.sel:
                    return
                global lastDraw
                try:
                    # 删除刚画完的图形，要不然鼠标移动的时候是黑乎乎的一片矩形
                    self.canvas.delete(lastDraw)
                except Exception as e:
                    pass
                lastDraw = self.canvas.create_rectangle(self.X.get(), self.Y.get(), event.x, event.y, outline='black')

            self.canvas.bind('<B1-Motion>', onLeftButtonMove)

            # 获取鼠标左键抬起的位置，保存区域截图
            def onLeftButtonUp(event):
                self.sel = False
                try:
                    self.canvas.delete(lastDraw)
                except Exception as e:
                    pass
                sleep(0.1)
                # 考虑鼠标左键从右下方按下而从左上方抬起的截图
                left, right = sorted([self.X.get(), event.x])
                top, bottom = sorted([self.Y.get(), event.y])
                pic = ImageGrab.grab((left + 1, top + 1, right, bottom))
                global paths
                paths = '测试.jpg'
                pic.save(paths)

                # 弹出保存截图对话框
                # fileName = tkinter.filedialog.asksaveasfilename(title='保存截图', filetypes=[('image', '*.jpg *.png')])
                # if fileName:
                #     pic.save(fileName)
                # 关闭当前窗口
                self.top.destroy()

            self.canvas.bind('<ButtonRelease-1>', onLeftButtonUp)
            # 让canvas充满窗口，并随窗口自动适应大小
            self.canvas.pack(fill=tkinter.BOTH, expand=tkinter.YES)

    # 开始截图
    def buttonCaptureClick():
        window.state('icon')  # 最小化主窗口
        sleep(0.2)
        filename = 'temp.png'
        # grab()方法默认对全屏幕进行截图
        im = ImageGrab.grab()
        im.save(filename)
        im.close()
        # 显示全屏幕截图
        w = MyCapture(filename)
        btn1.wait_window(w.top)  # 截图结束，恢复主窗口，并删除临时的全屏幕截图文件
        window.state('normal')  # 最大化主窗口
        remove(filename)  # 删除保存的全屏截图
        sb(paths)
        remove(paths)   # 删除保存的截图文件

    buttonCaptureClick()

window = Tk()
window.geometry('780x740+500+200')  # 窗口大小
window.title('中文图片文字识别')

taitouLabel = Label(window, text="请选择要识别的图片:", height=2, width=30, font=("Times", 20, "bold"), fg='red')
GunDongTiao = Scrollbar(window)  # 设置滑动块组件
confirmLabel = Text(window, height=20, width=55, font=("Times", 15, "bold"), fg='red', bg='#EEE5DE',
                    yscrollcommand=GunDongTiao.set)  # Listbox组件添加Scrollbar组件的set()方法

btn = Button(window, text="文件识别", command=xz, font=("Times", 15, "bold"))
btn1 = Button(window, text="截图识别", command=JieTu, font=("Times", 15, "bold"))
GunDongTiao.config(command=confirmLabel.yview)  # 设置Scrollbar组件的command选项为该组件的yview()方法

taitouLabel.grid(column=1)
btn.grid(row=0, column=3)
btn1.grid(row=1, column=3)
confirmLabel.grid(row=3, column=1, sticky=E, columnspan=3)
GunDongTiao.grid(row=3, column=5, sticky=N + S + W, )  # 设置垂直滚动条显示的位置

window.mainloop()


