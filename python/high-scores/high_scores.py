def latest(scores):
    # return (scores and scores[-1]) or None
    return scores[-1] if scores else None


def personal_best(scores):
    return max(scores) if scores else None


def personal_top_three(scores):
    return sorted(scores, reverse=True)[:3]
