import matplotlib.pyplot as plt
import patternMatching
import subprocess
import numpy as np
from sklearn.linear_model import LinearRegression

base_text = "AABA"
inc_text = base_text
pattern = "AABA"
num_test = 10


def do_prediction(x_axis, y_axis, text_length_2_predict=1000):
    """

    :param x_axis:
    :param y_axis:
    :param text_length_2_predict:
    :return:
    """
    X = np.array(x_axis).reshape(-1, 1).astype(np.float64)
    y = np.array(y_axis).reshape(-1, 1).astype(np.float64)
    reg = LinearRegression().fit(X, y)
    #print(reg.score(X, y))
    #print(reg.coef_)
    return reg.predict(np.array([[text_length_2_predict]]))


def plot(file_name, y_label, save):
    """

    :param file_name:
    :param y_label:
    :param save:
    :return:
    """
    f = open(file_name, "r")
    y_axis = []
    x_axis = []
    for j in range(num_test):
        y_axis.append(f.readline())
        x_axis.append(4 + j)
    plt.plot(x_axis, y_axis, 'r')
    plt.xlabel('Text length')
    plt.ylabel(y_label)

    if save:
        plt.savefig('test/'+y_label+'.png')
    plt.show(block=False)
    return x_axis, y_axis


if __name__ == '__main__':

    for i in range(num_test):
        patternMatching.search(pattern, inc_text, False, True)
        subprocess.run(["pattern_matching", str(i), "encr_batch"])
        inc_text = inc_text + "A"

    #read EncrTime and plot results
    x_axis, y_axis = plot('test/EncrTime.txt', 'Encryption Time (ms)', True)
    print(do_prediction(x_axis, y_axis, text_length_2_predict=1000))



