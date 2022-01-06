from dataclasses import dataclass
from typing import List, Optional

from dataclasses_json import dataclass_json


@dataclass_json
@dataclass
class Qa:
    question: str
    answer: str
    characteristic: Optional[List[str]] = None


@dataclass_json
@dataclass
class Text:
    raw_text: str
    characteristic: Optional[List[str]] = None
