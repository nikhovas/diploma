import ssl

import nltk
from string import punctuation
import gensim.downloader as api
import pymorphy2
from nltk.corpus import stopwords

from service_types import Text, Qa


class Model:
    def __init__(self):
        try:
            _create_unverified_https_context = ssl._create_unverified_context
        except AttributeError:
            pass
        else:
            ssl._create_default_https_context = _create_unverified_https_context
        nltk.download("stopwords")

        self.punc = punctuation + '–'
        self.russian_stopwords = stopwords.words("russian")
        self.model = api.load('word2vec-ruscorpora-300')
        self.morph = pymorphy2.MorphAnalyzer()
        self.tag_fix = {'ADJF': 'ADJ', 'ADJS': 'ADJ', 'INFN': 'VERB', 'PRED': 'ADV', 'ADVB': 'ADV', 'CONJ': 'SCONJ'}

    def pp(self, sentence):
        for c in self.punc:
            sentence = sentence.replace(c, ' ')
        result = []
        for w in sentence.lower().split():
            if w not in self.russian_stopwords:
                p = self.morph.parse(w)[0]
                nf = p.normal_form
                pos = p.tag.POS if p.tag.POS is not None else ''
                if pos in set(self.tag_fix.keys()):
                    pos = self.tag_fix[pos]
                result.append(nf + '_' + pos)
        return result

    def get_distance(self, first, second):
        denominator = 1 + 0.1 * len(set(first.characteristic) & set(second.characteristic))
        return self.model.wmdistance(first.characteristic, second.characteristic) / denominator

    def new_text(self, raw_text: str) -> Text:
        text = Text(raw_text=raw_text)
        text.characteristic = self.pp(raw_text)
        return text

    def new_qa(self, question: str, answer: str) -> Qa:
        qa = Qa(question=question, answer=answer)
        qa.characteristic = []
        qa.characteristic += self.pp(question)
        qa.characteristic += self.pp(answer)
        return qa
