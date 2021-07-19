import logging
import datetime
import json
from pythonjsonlogger import jsonlogger


class CustomJsonFormatter(jsonlogger.JsonFormatter):
    def add_fields(self, log_record, record, message_dict):
        super(CustomJsonFormatter, self).add_fields(log_record, record, message_dict)
        if not log_record.get('log_time'):
            # this doesn't use record.created, so it is slightly off
            now = datetime.datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%S.%fZ')
            log_record['log_time'] = now
        if log_record.get('level'):
            log_record['level'] = log_record['level'].upper()
        else:
            log_record['level'] = record.levelname
            
logger = logging.getLogger()
logger.setLevel(logging.DEBUG)
formatter = CustomJsonFormatter('%(log_time)s %(level)s %(name)s %(message)s')

stream_logHandler = logging.StreamHandler()
stream_logHandler.setLevel(logging.DEBUG)
stream_logHandler.setFormatter(formatter)

log_filename = datetime.datetime.now().strftime("%Y-%m-%d_%H_%M_%S.log")
file_logHandler = logging.FileHandler("data/ner-service/{}".format(log_filename))
file_logHandler.setLevel(logging.DEBUG)
file_logHandler.setFormatter(formatter)
logger.addHandler(stream_logHandler)
logger.addHandler(file_logHandler)
import time
st = time.time()
count = 0
while time.time() - st < 20000:
    payload = {"id": "60e55e32f321e00001de82c5", "data": "hello"}
    logger.info("model_predicts: {}".format(json.dumps({"user_id": "60e55e32f321e00001de82c5", "version_id": "60e56863b751572cbd49e0b1", "time": "{}".format(count)})))
    time.sleep(10)
    count=count+1
