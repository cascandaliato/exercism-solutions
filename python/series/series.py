def slices(series, length):
    if not series or length <= 0 or length > len(series):
        raise ValueError("Invalid input")
    return [series[i:(i+length)] for i in range(len(series)-length+1)]
