import matplotlib.pyplot as plt
from datetime import datetime
import matplotlib.dates as mdates
import numpy as np
import matplotlib as mpl
file = open('axis.txt')  
data = file.readlines() 
time = []  
mem_res_MB = []  

for num in data:
    t = datetime.strptime(num.split(' ')[0], '%H:%M:%S')
    time.append(t)
    mem_res_MB.append(float(num.split(' ')[1])/1024)

mpl.rcParams['font.size'] = 24

plt.figure(figsize=(20, 18))
plt.plot(time, mem_res_MB,color='blue', linewidth=0.2)

plt.gca().xaxis.set_major_formatter(mdates.DateFormatter("%H:%M:%S"))
plt.plot(time, mem_res_MB, zorder=1)
plt.xlabel("Time")#横坐标名字
plt.ylabel("Memory Usage(MB)")#纵坐标名字
plt.gcf().autofmt_xdate()
plt.yticks(range(int(min(mem_res_MB)), int(max(mem_res_MB)) + 1, 60))

peak_index = np.argmax(mem_res_MB)
peak_x_value = time[peak_index]
peak_y_value = mem_res_MB[peak_index]
max_x_pos = datetime.strptime("16:40:59", "%H:%M:%S")


min_index = np.argmin(mem_res_MB)
min_x_value = time[min_index]
min_y_value = mem_res_MB[min_index]
pos_x_min = datetime.strptime("15:56:02", "%H:%M:%S")
pos_y_min = float(28)

time_to_plot = datetime.strptime("16:22:59", "%H:%M:%S")
pos_to_plot = datetime.strptime("16:20:59", "%H:%M:%S")
value_to_plot = float(377980)/1024
pos_t_plot = value_to_plot - float(74)

plt.scatter(peak_x_value, peak_y_value, c='red', s=86, marker='o', label=f'max ({peak_x_value}, {peak_y_value})', zorder=3)
plt.annotate(f'Max Memory Usage: {peak_y_value}', (max_x_pos, peak_y_value + float(3)), textcoords="offset points", xytext=(0, 10), ha='center')
plt.scatter(min_x_value, min_y_value, c='red', s=86, marker='o', label=f'min ({min_x_value}, {min_y_value})', zorder=3)
plt.annotate(f'Minimal Memory Usage: {min_y_value}', (pos_x_min, pos_y_min), textcoords="offset points", xytext=(0, 10), ha='center')

plt.scatter(time_to_plot, value_to_plot, c='red', marker='o', s=86, label=f'16:22:59 ({time_to_plot}, {value_to_plot})', zorder=3)
plt.annotate(f'Memory Usage\nWhen Rule All Started: {value_to_plot}', (pos_to_plot, pos_t_plot), textcoords="offset points", xytext=(0, 10),  ha='left')


plt.savefig('visual_exp.svg', format='svg')