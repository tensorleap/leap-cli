from code_loader.contract.datasetclasses import PredictionTypeHandler
from code_loader.plot_functions.visualize import visualize
from leap_binder import *
from code_loader.inner_leap_binder.leapbinder_decorators import tensorleap_load_model, tensorleap_integration_test, tensorleap_status_table

prediction_type1 = PredictionTypeHandler('<place predictions name here>', "<place labels dict here>")

@tensorleap_load_model([prediction_type1])
def load_model():
    """
    Loads and returns the model for Tensorleap integration testing.

    Supports .h5 and .onnx models. Called once during the integration
    test to provide the model instance for inference.

    Example:
        model_path = 'models/model.onnx'
        sess = onnxruntime.InferenceSession(model_path)
        return sess

    Returns:
        Model object (e.g., onnxruntime.InferenceSession or tf.keras.Model).
    """
    pass


@tensorleap_integration_test()
def integration_test(idx, subset):
    """
    Defines and runs the local Tensorleap integration test.

    Simulates Tensorleap’s data flow by connecting encoders,
    model inference, loss, metrics, metadata, and visualizers.

    Example:
        image = input_encoder(idx, subset)
        gt = gt_encoder(idx, subset)
        model = load_model()
        y_pred = model([image])
        loss = custom_loss(gt, y_pred)

    Args:
        idx (int): Sample index to test.
        subset (PreprocessResponse): Dataset partition.

    Returns:
        None: Used for validation and debugging of integration flow.
    """

    # image = input_encoder(idx, subset)
    # gt = gt_encoder(idx, subset)
    # model = load_model()
    # y_pred = model([image])
    # loss = custom_loss(gt, y_pred)
    # metric=custom_metric(gt, y_pred)
    # metadata=custom_metadata(idx, subset)
    # visualizer_obj=visualize(image)
    # visualize(visualizer_obj)



if __name__ == '__main__':
    tensorleap_status_table() # Checks users implementation status
    integration_test(0, preprocess_func_leap()[0]) # DO NOT MOVE THIS OUTSIDE OF __name__==__main__ block!












