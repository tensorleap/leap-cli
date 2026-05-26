from typing import Union

import numpy as np
from code_loader.contract.datasetclasses import PreprocessResponse, PredictionTypeHandler
from code_loader.inner_leap_binder.leapbinder_decorators import (
    tensorleap_load_model,
    tensorleap_integration_test,
)

LABELS = []  # TODO: define your class labels

prediction_type = PredictionTypeHandler('output', LABELS, channel_dim=-1)


@tensorleap_load_model([prediction_type])
def load_model():
    pass  # TODO: load and return your model


@tensorleap_integration_test()
def integration_test(idx: Union[int, str], subset: PreprocessResponse) -> None:
    model = load_model()

    # inputs & gt
    # predictions
    # visualizers
    # metrics & losses
    # metadata


if __name__ == '__main__':
    pass  # TODO: call your preprocess function and run integration_test
