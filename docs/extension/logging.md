# Logrus

In order to use [logrus] as your logger GoEngine provides a wrapper for both `*logrus.Logger` and `*logrus.Entry`. 

```golang
import (
	"github.com/vimeda/goengine"
	goengineLogger "github.com/vimeda/goengine/extension/logrus"
	"github.com/sirupsen/logrus"
)

var logger goengine.Logger
logger = goengineLogger.Wrap(s.Logger)
logger = goengineLogger.WrapEntry(s.Logger.WithField("source": "goengine"))
```

# Zap

In order to use [zap] as your logger GoEngine provides a wrapper for `*zap.Logger`. 

```golang
import (
	"github.com/vimeda/goengine"
	goengineLogger "github.com/vimeda/goengine/extension/zap"
	"go.uber.org/zap"
)

var logger goengine.Logger
logger = goengineLogger.Wrap(zap.NewNop())
```


[logrus]: https://github.com/sirupsen/logrus
[zap]: https://github.com/uber-go/zap/
