import matplotlib.pyplot as plt
import patternMatching
import subprocess
import numpy as np
from sklearn.linear_model import LinearRegression

TEXT_LENGTH_2_PREDICT = 1000

BASE_TEXT = "AABA"
inc_text = BASE_TEXT
PATTERN = "AABA"
NUMMBER_OF_TESTS = 40


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


def plot(x_axis, y_axis, y_label, save):
    '''

    :param x_axis:
    :param y_axis:
    :param y_label:
    :param save:
    :return:
    '''
    plt.plot(x_axis, y_axis, 'r')
    plt.xlabel('Text length')
    plt.ylabel(y_label)

    if save:
        plt.savefig('test/'+y_label+'.png')
    plt.show()


def get_axis(file_name):
    """

    :param file_name:
    :return:
    """
    f = open(file_name, "r")
    y_axis = []
    x_axis = []
    for j in range(NUMMBER_OF_TESTS):
        y_axis.append(f.readline())
        x_axis.append(4 + j)
    return x_axis, y_axis


if __name__ == '__main__':

    for i in range(NUMMBER_OF_TESTS):
        patternMatching.search(PATTERN, inc_text, False, True)
        subprocess.run(["pattern_matching", str(i), "all"])
        inc_text = inc_text + "A"

    #read EncrTime, plot results and do prediction
    x_axis_encr, y_axis_encr = get_axis('test/EncrTime.txt')
    print("Encryption time for text of length "+str(TEXT_LENGTH_2_PREDICT)+": "+str(do_prediction(x_axis_encr, y_axis_encr, TEXT_LENGTH_2_PREDICT)))
    plot(x_axis_encr, y_axis_encr, 'Encryption Time (ms)', True)

    #read DecrTime, plot results and do prediction
    x_axis_decr, y_axis_decr = get_axis('test/DecrTime.txt')
    print("Decryption time for text of length "+str(TEXT_LENGTH_2_PREDICT)+": "+str(do_prediction(x_axis_decr, y_axis_decr, TEXT_LENGTH_2_PREDICT)))
    plot(x_axis_decr, y_axis_decr, 'Decryption Time (ms)', True)



